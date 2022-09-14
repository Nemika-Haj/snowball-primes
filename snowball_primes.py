import math
 
def is_prime(n: int) -> bool:
    for i in range(2, int(math.sqrt(n)) + 1):
        if (n % i == 0):
            return False
    return True

if __name__ == "__main__":
    try:
        snowball = int(input("Enter the first digit of the snowball: "))
        #if not is_prime(snowball): raise ValueError

        snowballs = [snowball]

        while all([is_prime(snowball) for snowball in snowballs]):
            new_snowballs = []

            for snowball in snowballs:
                for i in range(1, 10, 2):
                    if is_prime(snowball*10+i):
                        new_snowballs.append(snowball*10+i)
            

            if not new_snowballs: break

            snowballs = new_snowballs

        
        print(snowballs[-1], "is the biggest snowball prime number from that starting point!")
            

    except ValueError:
        print("A prime integer is required. Try again.")