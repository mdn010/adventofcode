require 'set'

required_fields = Set['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']

passports = File.read('input.txt').split(/\n{2,}/)

answer = 0

passports.each do |passport|
	set = Set[]
	passport.split.each do |record|
		parsed = record.split(':')
		field = parsed.first
		value = parsed[1]
		case field
		when 'byr'
			set.add(field) if (1920..2002).include?(value.to_i)
		when 'iyr'
			set.add(field) if (2010..2020).include?(value.to_i)
		when 'eyr'
			set.add(field) if (2020..2030).include?(value.to_i)
		when 'hgt'
			height = value.chomp('cm').chomp('in')
			if value.include?('cm')
				set.add(field) if (150..193).include?(value.to_i)
			elsif value.include?('in')
				set.add(field) if (59..76).include?(value.to_i)				
			end
		when 'hcl'
			set.add(field) if value.match?(/^#[0-9a-f]{6}$/)
		when 'ecl'
			set.add(field) if ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].include?(value)
		when 'pid'
			set.add(field) if value.match?(/^\d{9}$/)
		end
	end
	answer += 1 if (required_fields - set).empty?
end
