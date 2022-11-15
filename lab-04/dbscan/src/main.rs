extern crate image;
extern crate rand;

use image::GenericImageView;
use image::ImageBuffer;
use rand::Rng;
use std::cmp;


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



fn dbscan(points: Vec<Vec<bool>>, min_ptx: usize, mut cluster_count: u32, width: u32, height: u32, eps: f64) -> u32 {

    let mut imgbuf = image::ImageBuffer::new(width, height);
    for (_, _, pixel) in imgbuf.enumerate_pixels_mut() {
        *pixel = image::Rgb([255, 255, 255]);
    }

    let mut current_point = points.clone();
    let mut current_color = get_random_color();
    println!("{}", points[0][0]);

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
                    
                    regionquery(&points, min_ptx, &mut v, &mut neighbor_count, p, eps);

                    if neighbor_count >= min_ptx {
                        //if neighbor_count_check was hold
                        neighbor_count_check += 1;
                        *imgbuf.get_pixel_mut(p[1] as u32, p[0] as u32) = image::Rgb(current_color);
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
    imgbuf.save("./data/dbscan.png").unwrap();
    cluster_count
}


fn main() {

    let eps = 2.0;
    let min_ptx = 2;
    
    let img = image::open("./data/me.bmp").unwrap();
    let target = [0, 0, 0, 255];
    let mut points = Vec::<Vec<bool>>::new();
    let (width, height) = img.dimensions();
    for h in 0..height {
        let mut _p = Vec::<bool>::new();
        for w in 0..width {
            _p.push(&img.get_pixel(w, h).0 == &target);
        }
        points.push(_p);
    }
    
    let mut cluster_count = 0;

    
    cluster_count = dbscan(points, min_ptx, cluster_count, width, height, eps);
    println!("cluster {}", cluster_count);
}
