use std::fs;

fn main() {
    let contents = fs::read_to_string("data/input.txt")
        .expect("Should have been able to read the file");
    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();
    for row in contents.split("\n") {
        if row == "" {
            continue;
        }
        let row_split: Vec<&str> = row.split_whitespace().collect();
        match row_split[0].parse::<i32>() {
            Ok(n) => {
                left.push(n);
            }
            Err(e) => {
                panic!("{}", e)
            },
        }
        match row_split[1].parse::<i32>() {
            Ok(n) => {
                right.push(n);
            }
            Err(e) => {
                panic!("{}", e)
            },
        }
        
    }
    left.sort();
    right.sort();
    let result = left.iter().zip(right.iter()).map(|x| (x.0 - x.1).abs()).reduce(|acc, e| acc + e).unwrap();
    println!("{}", result);
}
