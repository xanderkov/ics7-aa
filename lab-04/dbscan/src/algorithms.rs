extern crate image;
extern crate rand;


use crossbeam::thread as cthread;
use std::sync::{ Arc, Mutex };
use image::GenericImageView;
use rand::Rng;
use std::cmp;
use std::time::Instant;
type MultFnPtr = fn(&Vec<Vec<bool>>, usize, u32, u32, f64) -> u32;

pub static MULTS_ARRAY: [MultFnPtr; 2] = [dbscan, dbscan_p];
pub static MULTS_DESCRIPTIONS: [&str; 2] = ["Простой dbscan", "Параллельный dbscan"];
const NUMBER_OF_THREADS: usize = 8;

fn get_random_color() -> [u8; 3] {
    let mut rng = rand::thread_rng();
    [rng.gen(), rng.gen(), rng.gen()]
}

fn get_eculid_distance(start_x: i32, start_y: i32, end_x: i32, end_y: i32) -> f64 {
    return (((start_x - end_x).pow(2) + (start_y - end_y).pow(2)) as f64).sqrt();
}

fn regionquery(points: &Vec<Vec<bool>>, min_ptx: usize, v: &mut Vec::<[usize; 2]>, neighbor_count: &mut usize, p: [usize; 2], eps: f64)  {
    let start_x = cmp::max(0, p[0] - min_ptx);
    let start_y = cmp::max(0, p[1] - min_ptx);
    let end_x = cmp::min(points.len(), p[0] + min_ptx + 1);
    let end_y = cmp::min(points[0].len(), p[1] + min_ptx + 1);
    for _i in start_x..end_x {
        //check how many neighbors in the distance
        for _j in start_y..end_y {
            let distance = get_eculid_distance(_i as i32, _j as i32, p[0] as i32, p[1] as i32);
            if distance <= eps && points[_i][_j] {
                *neighbor_count += 1;
                v.push([_i, _j]);
            }
        }
    }
}



pub fn dbscan(points: &Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64) -> u32 {

    let mut cluster_count = 0;
    let mut imgbuf = image::ImageBuffer::new(width, height);
    for (_, _, pixel) in imgbuf.enumerate_pixels_mut() {
        *pixel = image::Rgb([255, 255, 255]);
    }

    let mut current_point = points.clone();
    let mut current_color = get_random_color();

    for i in 0..points.len() {
        for j in 0..points[i].len() {
            if current_point[i][j] {

                let mut v = Vec::<[usize; 2]>::new();
                v.push([i, j]);
                let mut neighbor_count_check = 0;
                while !v.is_empty() {
                    let p = v.pop().unwrap();
                    if !current_point[p[0]][p[1]] {
                        continue;
                    }
                    current_point[p[0]][p[1]] = false;
                    
                    let mut neighbor_count = 0;
                    
                    regionquery(points, min_ptx, &mut v, &mut neighbor_count, p, eps);

                    if neighbor_count >= min_ptx {
                        neighbor_count_check += 1;
                        // *imgbuf.get_pixel_mut(p[1] as u32, p[0] as u32) = image::Rgb(current_color);
                    }
                }
                if neighbor_count_check > 0 {
                    cluster_count += 1;
                    current_color = get_random_color();
                }
                current_point[i][j] = false;
            }
        }
    }
    //imgbuf.save("./data/dbscan.png").unwrap();
    cluster_count
}


pub fn parallel_for(points: &Vec<Vec<bool>>, 
    min_ptx: usize, eps: f64, 
    range: std::ops::Range<usize>, 
    guard_copy: Arc<Mutex<Vec<Vec<bool>>>>, 
    n: Arc<Mutex<u32>>) 
{
    let mut cluster_count = n.lock().unwrap();

    let mut current_point = guard_copy.lock().unwrap();
    let mut current_color = get_random_color();

    for i in range {
        for j in 0..points[i].len() {
            if current_point[i][j] {

                let mut v = Vec::<[usize; 2]>::new();
                v.push([i, j]);
                let mut neighbor_count_check = 0;
                while !v.is_empty() {
                    let p = v.pop().unwrap();
                    if !current_point[p[0]][p[1]] {
                        continue;
                    }
                    current_point[p[0]][p[1]] = false;
                    
                    let mut neighbor_count = 0;
                    
                    regionquery(points, min_ptx, &mut v, &mut neighbor_count, p, eps);

                    if neighbor_count >= min_ptx {
                        //if neighbor_count_check was hold
                        neighbor_count_check += 1;
                    }
                }
                if neighbor_count_check > 0 {
                    *cluster_count += 1;
                    current_color = get_random_color();
                }
                current_point[i][j] = false;
            }
        }
    }

}

pub fn dbscan_parallel(points: &Vec<Vec<bool>>, min_ptx: usize, eps: f64, nofth: usize) -> u32 {

    let counter = Arc::new(Mutex::new(0));
    let current_point = Arc::new(Mutex::new(points.clone()));
    let size = points.len() / (nofth + 1);

    cthread::scope(|s| {
        let mut threads = Vec::with_capacity(nofth);

        for i in 0..nofth {
            let range = (i * size)..((i + 1) * size);
            let guard_copy = current_point.clone();
            let counter_copy = counter.clone();
            threads.push(s.spawn(move |_| parallel_for(points, min_ptx, eps, range, guard_copy, counter_copy)));
        }

        let guard_copy = current_point.clone();
        let range = (nofth * size)..points.len();
        let counter_copy = counter.clone();

        parallel_for(points, min_ptx, eps, range, guard_copy, counter_copy);
        for th in threads {
            th.join().unwrap();
        }

        let counter_copy = counter.clone();
        let cluster_count = counter_copy.lock().unwrap();
        *cluster_count
    }).unwrap()

    
}

pub fn dbscan_p(points: &Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64) -> u32 {
    let res = dbscan_parallel(points, min_ptx, eps, NUMBER_OF_THREADS - 1);
    res
}

pub fn run_tests(points: Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64) {
    for j in 1..5 {
        println!("Количество замеров: {} \n", j * 100);
        for (algorithm, description) in MULTS_ARRAY.iter().zip(MULTS_DESCRIPTIONS.iter()) {
            let time = Instant::now();
            let mut result = 0;
            for i in 0..j * 100 {
                // println!("Замер №{}", i + 1);
                result = algorithm(&points, min_ptx, width, height, eps);
            }

            
            let time = time.elapsed().as_nanos();
            println!("Алгоритм: {}", description);
            println!("Результат: {} ", result);
            println!("Время: {} нс \n", time);
        }
        
    }
}
