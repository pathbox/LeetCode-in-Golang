def length_of_longest_substring(s)
  return 0 if s.empty?
  return 1 if s.length == 1

  max_length = 0
  arr = s.chars
  len = s.length

  s.length.times do |index| # 从头到尾遍历字符串
    str = '' # 符合条件的子串
    (index..len-1).each do |i| # 每次遍历一个子串，并寻找复合条件的子串
      if str.include? arr[i] # 有重复字符，则退出遍历子串，得到当前的str
        break
      else
        str += arr[i]
      end
    end
    max_length = [max_length, str.length].max
  end
  max_length
end