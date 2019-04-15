class Solution(object):
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        constant = [
            ["","I","II","III","IV","V","VI","VII","VIII","IX"],
            ["","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"],
            ["","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"],
            ["","M","MM","MMM"]
        ]
        
        roman = list()
        roman.append(constant[3][num / 1000 % 10])
        roman.append(constant[2][num / 100 % 10])
        roman.append(constant[1][num / 10 % 10])
        roman.append(constant[0][num % 10])
        
        return "".join(roman)