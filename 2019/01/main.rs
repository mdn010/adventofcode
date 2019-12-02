fn main() {
    let input = include_str!("input.txt");

    let mut _iter = input
        .split(|c| c == '\n' || c == ' ')
        .filter_map(|l| l.parse::<usize>().ok());

    fn meta_sum(iter: &mut impl Iterator<Item = usize>) -> usize {
        iter.map(|m| m/3).map(|m| m-2).sum::<usize>()
    };

    println!("Part 1 : {}", meta_sum(&mut _iter));

    let mut part2 = input
        .split(|c| c == '\n' || c == ' ')
        .filter_map(|l| l.parse::<usize>().ok());

    fn miracle_sum(iter: &mut impl Iterator<Item = usize>) -> usize {
        iter.map(|m| m/3).map(|m| m - 2).map(|m| m + ssum(m)).sum::<usize>()
    };

    fn ssum(m: usize) -> usize {
        let xs = vec![m];
        xs.iter().map(|m| m/3).filter(|m| m >= &"2".parse::<usize>().unwrap()).map(|m| m - 2).map(|m| m + ssum(m)).sum::<usize>()
    };

    println!("Part 2 : {}", miracle_sum(&mut part2));
}