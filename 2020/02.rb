valids0 = File.read('input.txt').split("\n").select do |line|
	parts = line.split

	policy_count = parts[0].split('-')
	at_least = policy_count[0].to_i
	at_most = policy_count[1].to_i

	char = parts[1].chomp(':')

	password = parts[2]

	password.count(char).between?(at_least, at_most)
end

answer0 = valids0.count

valids1 = File.read('input.txt').split("\n").select do |line|
	parts = line.split

	policy_index = parts[0].split('-')
	pos1 = policy_index[0].to_i
	pos2 = policy_index[1].to_i

	char = parts[1].chomp(':')

	password = parts[2]

	cond = [pos1, pos2].select do |pos|
		return false if pos == 0

		password[pos-1] == char
	end

	!cond.nil? && cond.count == 1
end

answer1 = valids1.count
