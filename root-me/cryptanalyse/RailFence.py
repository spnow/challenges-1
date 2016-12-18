# author: youenn.laborde

# Rail Fence est une permutation des lettres.
# Exemple : abcdef -> RailFence(2) -> acebdf

chal = 'Weiilglh tihn.v aSdieg nKaeln tiusc k"yF royn  tOhcet ocbheirc ktehne" .'
str1 = ""
str2 = ""
for i in range(0,len(chal), 2):
    str1 += chal[i]
    str2 += chal[i+1]

print str1+str2
