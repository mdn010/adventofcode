arr = File.read('input.txt').split.map(&:to_i)

answer0 = arr.combination(2).find { |a| a.reduce(0, :+) == 2020 }.reduce(:*)

answer1 = arr.combination(3).find { |a| a.reduce(0, :+) == 2020 }.reduce(:*)
