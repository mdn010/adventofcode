use std::collections::HashMap;

fn main () {
    let beg = 125730;
    let end = 579381;
    let mut count = 0;

    (beg..end+1).for_each(|n|{
        let mut hm:HashMap<_,_> = HashMap::new();
        let mut no:HashMap<_,_> = HashMap::new();
        for (i, v) in n.to_string().chars().enumerate() {
            hm.entry(i).or_insert(v.to_digit(10).unwrap());
            *no.entry(v.to_digit(10).unwrap()).or_insert(0) += 1;
        }
        hm.get(&0).filter(|d| *d<=hm.get(&1).unwrap())
            .and_then(|_| hm.get(&1).filter(|d| *d<=hm.get(&2).unwrap()))
            .and_then(|_| hm.get(&2).filter(|d| *d<=hm.get(&3).unwrap()))
            .and_then(|_| hm.get(&3).filter(|d| *d<=hm.get(&4).unwrap()))
            .and_then(|_| hm.get(&4).filter(|d| *d<=hm.get(&5).unwrap()))
            // Part 2
            .and_then(|_| hm.get(&0).filter(|d| *d==hm.get(&1).unwrap() && *no.get(d).unwrap() == 2)
                .or(hm.get(&1).filter(|d| *d==hm.get(&2).unwrap() && *no.get(d).unwrap() == 2))
                .or(hm.get(&2).filter(|d| *d==hm.get(&3).unwrap() && *no.get(d).unwrap() == 2))
                .or(hm.get(&3).filter(|d| *d==hm.get(&4).unwrap() && *no.get(d).unwrap() == 2))
                .or(hm.get(&4).filter(|d| *d==hm.get(&5).unwrap() && *no.get(d).unwrap() == 2)))
            .map(|_| count += 1);
    });

    println!("{:#?}", count);
}
