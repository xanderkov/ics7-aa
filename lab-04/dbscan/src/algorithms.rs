extern crate image;
extern crate rand;



use crossbeam::thread as cthread;
use std::sync::{ Arc, Mutex };
use super::*;
use image::GenericImageView;
use rand::Rng;
use std::cmp;
use std::time::Instant;
type MultFnPtr = fn(&Vec<Vec<bool>>, usize, u32, u32, f64) -> u32;

pub static MULTS_ARRAY: [MultFnPtr; 1] = [dbscan];
pub static MULTS_DESCRIPTIONS: [&str; 1] = ["Простой dbscan"];

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


pub fn parallel_for(points: &Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64, nofth: usize) -> u32 {

    cthread::scope(|s| {
        let mut threads = Vec::with_capacity(nofth);

        let size = points.len() / (nofth + 1);
        let points_guard = Arc::new(Mutex::new(points));

        for i in 0..nofth {
            let range = (i * size)..((i + 1) * size);
            let guard_copy = points_guard.clone();
            let precopy = &precomputed;
            threads.push(s.spawn(move |_| multiplication_in_range(&precopy, guard_copy, &m1, &m2, range)));
        }

        let range = (nofth * size)..points.len();
        let guard_copy = points_guard.clone();
        multiplication_in_range(&precomputed, guard_copy, &m1, &m2, range);
        for th in threads {
            th.join().unwrap();
        }

        let result = Arc::try_unwrap(points_guard).unwrap().into_inner().unwrap();

        result
    }).unwrap()

}

pub fn dbscan_parallel(points: &Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64) -> u32 {
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
                        //if neighbor_count_check was hold
                        neighbor_count_check += 1;
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
    cluster_count
}

pub fn run_tests(points: Vec<Vec<bool>>, min_ptx: usize, width: u32, height: u32, eps: f64) {
    for (algorithm, description) in MULTS_ARRAY.iter().zip(MULTS_DESCRIPTIONS.iter()) {
        let time = Instant::now();
        
        for i in 0..5 {
            println!("Замер №{}", i + 1);
            let result = dbscan(&points, min_ptx, width, height, eps);
        }
        
        let time = time.elapsed().as_nanos();
        println!("{}: ", description);
        println!("\nВремя: {} нс", time);
    }
}
