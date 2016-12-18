#BMP files have the same kind of headers. I took a blanc bmp file, and xor it with ch3 file.
#Because they have the same headers, we can see the key in the begining.

#This key was guess by the first function key_recover
key = [ord('f'),ord('a'),ord('l'),ord('l'),ord('e'),ord('n')]

def key_recover():
    '''
    Recover the key
    '''
    #The File From the Challenge
    file = open('ch3.bmp','rb')
    dataY = bytearray(file.read())
    file.close()

    #A blank bmp Image to compare headers
    file = open('blank.bmp','rb')
    dataX = bytearray(file.read())
    file.close()

    #XOR both image to recover key
    print "X \t Y \t X^Y \t ASCII"
    print "----------------------------------"
    for i in range(0,48):
        print dataX[i], " \t", dataY[i], " \t", (dataX[i]^dataY[i]), "  \t", chr(dataX[i]^dataY[i])

def decode():
    '''
    Decode encrypted bmp file
    '''
    #The File from the Challenge
    file = open('ch3.bmp','rb')
    data = bytearray(file.read())
    file.close()

    #Decode each element
    for i in range(0,len(data)):
        data[i] = data[i] ^ key[i%len(key)]

    #Save into a file
    file = open('res.bmp','wb')
    file.write(data)
    file.close()

decode()
