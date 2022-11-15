extern crate image;
extern crate rand;

use image::GenericImageView;
mod algorithms;


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

    algorithms::run_tests(points, min_ptx, width, height, eps);
    //let cluster_count = algorithms::dbscan(points, min_ptx, width, height, eps);
    //println!("cluster {}", cluster_count);
}
