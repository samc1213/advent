use std::fs;

fn is_safe(row: &Vec<i32>) -> bool {
    let at_least_one_at_most_three = row.windows(2).all(|pair| (pair[0] - pair[1]).abs() >= 1 && (pair[0] - pair[1]).abs() <= 3);
    let all_increasing = row.windows(2).all(|pair| pair[0] < pair[1]);
    let all_decreasing = row.windows(2).all(|pair| pair[0] > pair[1]);
    return at_least_one_at_most_three && (all_increasing || all_decreasing);
}

fn main() {
    let contents = fs::read_to_string("data/input.txt")
        .expect("Should have been able to read the file");
    let mut total = 0;
    for row in contents.split("\n") {
        if row == "" {
            continue;
        }
        let row_split: Vec<i32> = row
            .split_whitespace()
            .map(|s| s.parse::<i32>().expect("Invalid number"))
            .collect();
        if is_safe(&row_split) {
            total += 1;
        }
    }
    println!("{}", total);
}
