extern crate image;
extern crate rand;

use image::GenericImageView;
mod utils;

fn main() {

    let eps = 2.0;
    let min_ptx = 2;
    
    let img = image::open("./data/1kme.bmp").unwrap();
    let target = [0, 0, 0, 255];
    let mut points = Vec::<Vec<bool>>::new();
    let (width, height) = (500, 500);
    for h in 0..height {
        let mut _p = Vec::<bool>::new();
        for w in 0..width {
            _p.push(&img.get_pixel(w, h).0 == &target);
        }
        points.push(_p);
    }

    let mut imgbuf = image::RgbImage::new(width, height);
    for (_, _, pixel) in imgbuf.enumerate_pixels_mut() {
        *pixel = image::Rgb([255, 255, 255]);
    }
    //algorithms::dbscan(&points, min_ptx, width,height,eps, &imgbuf);
    utils::run_tests(points, min_ptx, eps, imgbuf);
    //let cluster_count = algorithms::dbscan(points, min_ptx, width, height, eps);
    //println!("cluster {}", cluster_count);
}
