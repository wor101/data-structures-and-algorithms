bigArray = [1, 2, 3, [4, 5, 6], 7, [8, [9, 10, 11], [12, 13, 14]]]

def printNumbers(array)
  array.each do |element|
    if element.kind_of?(Array)
      printNumbers(element)
    elsif 
      puts "Element: " + element.to_s
    end
  end
end

printNumbers(bigArray)