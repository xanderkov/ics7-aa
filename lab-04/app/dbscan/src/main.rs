extern crate image;
extern crate rand;

use image::GenericImageView;
use image::ImageBuffer;
use rand::Rng;


fn main() {

    let img = image::open("./data/dbscan.png");
    let eps = 0.2;
    let minpit = 2;
    let (width, height) = img.dimensions();

    println!("{}, {}", eps, minpit);
    println!("{}, {}", width, height);
}
