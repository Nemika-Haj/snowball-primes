def is_prime(n)
    for i in 2...Math.sqrt(n).to_i+1
        if n % i == 0
            return false
        end
    end
    return true
end

print "Enter the first digit of the snowball: "
snowball = gets.chomp().to_i

snowballs = [snowball]

while true
    new_snowballs = []
    
    snowballs.each { |snowball|
        for i in (1..10).step(2)
            if is_prime(snowball*10 + i)
                new_snowballs << snowball*10 + i
            end
        end
    }
    
    if new_snowballs.empty?
        break
    end
    
    snowballs = new_snowballs
end

puts "#{snowballs[-1]} is the biggest snowball prime number from that starting point!"

