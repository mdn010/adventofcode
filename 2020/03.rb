grid = File.read('input.txt').split

offsets = [[1, 1], [1, 3], [1, 5], [1, 7], [2, 1]]
row_length = grid[0].size
counter = []

offsets.each do |offset|
	trees = 0
	col = 0
	row_offset, col_offset = offset[0], offset[1]
	(row_offset..grid.length - 1).step(row_offset).each do |row|
		col = (col + col_offset) % row_length
		trees += 1 if grid[row][col] == '#'
	end
	counter.push(trees)
end

answer = counter.reduce(:*)
