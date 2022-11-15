pub fn run_tests(m1: &[Vec<MatInner>], m2: &[Vec<MatInner>]) {
    for (algorithm, description) in MULTS_ARRAY.iter().zip(MULTS_DESCRIPTIONS.iter()) {
        let time = Instant::now();
        let result = algorithm(m1, m2);
        let time = time.elapsed().as_nanos();

        print!("{}: ", description);
        print_matrix(&result);
        println!("\nВремя: {} нс", time);
    }
}