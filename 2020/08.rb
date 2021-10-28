require 'set'

input = File.read("input.txt").split("\n")

def update(input, i)
	command = input[i].split
	operator = command[0]
	value = command[1].to_i
	case operator
	when "jmp"
		operator = "nop"
	when "nop"
		operator = "jmp"
	end
	input[i] = "#{operator} #{sprintf('%+d', value)}"
	return operator, value, input
end

def run(input, i, acc, set, commands)
	if set.include?(i)
		return acc, false
	else
		set.add(i)
		commands.push(i)
	end
	command = input[i].split
	operator = command[0]
	value = command[1].to_i
	case operator
	when "jmp"
		i = i + value
	when "acc"
		acc += value
		i += 1
	else
		i += 1
	end
	if i >= input.length
		return acc, true
	elsif !(0...input.length).include?(i)
		return acc, false
	else
		run(input, i, acc, set, commands)
	end
end

commands = []
ans0, success = run(input, 0, 0, Set[], commands)

(0...commands.length).reverse_each do |i|
	set = Set[]
	acc = 0
	operator, value, mod = update(input.clone, commands[i])
	acc, success = run(mod, 0, 0, Set[], [])
	if success
		puts acc
		break
	end
end
