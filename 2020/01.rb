arr = File.read('input.txt').split.map(&:to_i)

answer = arr.combination(2).find { |a| a.reduce(0, :+) == 2020 }.reduce(:*)

puts answer
