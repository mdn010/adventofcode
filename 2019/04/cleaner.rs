fn is_ramp(password: &Vec<u32>) -> bool { (0..5).all(|i| password[i] <= password[i+1]) }

fn has_doubles(password: &Vec<u32>) -> bool { (0..5).any(|i| password[i] == password[i+1]) }

fn has_doubles_only(password: &Vec<u32>) -> bool {
    (0..5).filter(|&i| password[i] == password[i+1])
        .any(|i| password.iter().filter(|&n| *n == password[i]).count() == 2)
}

fn main () {
    let range1 = (125730..579381+1).map(|i| i.to_string().chars().map(|c| c.to_digit(10).unwrap()).collect());
    let range2 = range1.clone();

    let part1 = range1.filter(is_ramp).filter(has_doubles).count();
    let part2 = range2.filter(is_ramp).filter(has_doubles_only).count();
    
    println!("{:#?}", part1);
    println!("{:#?}", part2);
}
