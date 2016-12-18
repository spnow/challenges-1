# cryptopals crypto challenges

My work in Golang for the [cryptopals crypto challenges](http://cryptopals.com). This is one big **spoiler**, do not watch it if you want to do the challenges by yourself.

- [x] **Set 1 : Basics**
  * [x] [01-Convert hex to base64](http://cryptopals.com/sets/1/challenges/1)
  * [x] [02-Fixed XOR](http://cryptopals.com/sets/1/challenges/2)
  * [x] [03-Single-byte XOR cipher](http://cryptopals.com/sets/1/challenges/3)
  * [x] [04-Detect single-character XOR](http://cryptopals.com/sets/1/challenges/4)
  * [x] [05-Implement repeating-key XOR](http://cryptopals.com/sets/1/challenges/5)
  * [x] [06-Break repeating-key XOR](http://cryptopals.com/sets/1/challenges/6)
  * [x] [07-AES in ECB mode](http://cryptopals.com/sets/1/challenges/7)
  * [x] [08-Detect AES in ECB mode](http://cryptopals.com/sets/1/challenges/8)
- [ ] **Set 2 : Block crypto**
  * [ ] [09-Implements PKCS#7 padding](https://cryptopals.com/sets/2/challenges/9)
  * [ ] [10-Implements CBC mode](https://cryptopals.com/sets/2/challenges/10)
  * [ ] [11-AnECb/CBC detection](https://cryptopals.com/sets/2/challenges/11)
  * [ ] [12-Byte-at-a-time ECB decryption Simple](https://cryptopals.com/sets/2/challenges/12)
  * [ ] [13-ECB cut-and-paste](https://cryptopals.com/sets/2/challenges/13)
  * [ ] [14-Byte-at-a-time ECB decryption Harder](https://cryptopals.com/sets/2/challenges/14)
  * [ ] [15-PKCS#7 âdding validation](https://cryptopals.com/sets/2/challenges/15)
  * [ ] [16-CBC bitflipping attacks](https://cryptopals.com/sets/2/challenges/16)
- [ ] **Set 3 : Block & stream crypto**
 * [ ] [17-The CBC padding oracle](https://cryptopals.com/sets/3/challenges/17)
 * [ ] [18-Implement CTR, the stream cipher mode](https://cryptopals.com/sets/3/challenges/18)
 * [ ] [19-Break fixed-nonce CTR mode using substitutions](https://cryptopals.com/sets/3/challenges/19)
 * [ ] [20-Break fixed-nonce CTR statistically](https://cryptopals.com/sets/3/challenges/20)
 * [ ] [21-Implement the MT19337 Mersenne Twister RNG](https://cryptopals.com/sets/3/challenges/21)
 * [ ] [22-Crack an MT19337 seed](https://cryptopals.com/sets/3/challenges/22)
 * [ ] [23-Clone an MT19337 RNG from its output](https://cryptopals.com/sets/3/challenges/23)
 * [ ] [24-Create the MT19337 stream wipher and break it](https://cryptopals.com/sets/3/challenges/24)
- [ ] **Set 4 : Stream crypto and randomness**
 * [ ] [25-Break random access read/write AES/CTR](https://cryptopals.com/sets/4/challenges/25)
 * [ ] [26-CTR bitflipping](https://cryptopals.com/sets/4/challenges/26)
 * [ ] [27-Recover the key from CBC with IV=Key](https://cryptopals.com/sets/4/challenges/27)
 * [ ] [28-Implement a SHA-1 keyed MAC](https://cryptopals.com/sets/4/challenges/28)
 * [ ] [29-Break a SHA1 keyed MAC using length extension](https://cryptopals.com/sets/4/challenges/29)
 * [ ] [30-Break a MD4 keyed MAC using length extension](https://cryptopals.com/sets/4/challenges/30)
 * [ ] [31-Implement and break HMAC-SHA1 with an artificial timing leak](https://cryptopals.com/sets/4/challenges/31)
 * [ ] [32-Break HMAC-SHA1 with a slightly less artificial timing leak](https://cryptopals.com/sets/4/challenges/32)
- [ ] **Set 5 : Diffie-Hellman and friends**
 * [ ] [33-Implement Diffie-Hellman](https://cryptopals.com/sets/5/challenges/33)
 * [ ] [34-Implement a MITM key-fixing attack on Diffie-Hellman with parameter injection](https://cryptopals.com/sets/5/challenges/34)
 * [ ] [35-Implement DH with negotiated groups, and break with malicious "g" parameters](https://cryptopals.com/sets/5/challenges/35)
 * [ ] [36-Implement Secure Remote Password (SRP)](https://cryptopals.com/sets/5/challenges/36)
 * [ ] [37-Break SRP with a zero key](https://cryptopals.com/sets/5/challenges/37)
 * [ ] [38-Offline dictionary attack on simplified SRP](https://cryptopals.com/sets/5/challenges/38)
 * [ ] [39-Implement RSA](https://cryptopals.com/sets/5/challenges/39)
 * [ ] [40-Implement an E=3 RSA Broadcast attack](https://cryptopals.com/sets/5/challenges/40)
- [ ] **Set 6 : RSA and DSA**
 * [ ] [41-Implement unpadded message recovery oracle](https://cryptopals.com/sets/6/challenges/41)
 * [ ] [42-Bleichenbacher's e=3 RSA Attack](https://cryptopals.com/sets/6/challenges/42)
 * [ ] [43-DSA key recovery from nonce](https://cryptopals.com/sets/6/challenges/43)
 * [ ] [44-DSA nonce recovery from repeated nonce](https://cryptopals.com/sets/6/challenges/44)
 * [ ] [45-DSA parameter tampering](https://cryptopals.com/sets/6/challenges/45)
 * [ ] [46-RSA parity oracle](https://cryptopals.com/sets/6/challenges/46)
 * [ ] [47-Bleichenbacher's PKCS 1.5 Padding Oracle (Simple Case)](https://cryptopals.com/sets/6/challenges/47)
 * [ ] [48-Bleichenbacher's PKCS 1.5 Padding Oracle (Complete Case)](https://cryptopals.com/sets/6/challenges/48)
- [ ] **Set 7 : Hashes**
 * [ ] [49-CBC-MAC Message Forgery](https://cryptopals.com/sets/7/challenges/49)
 * [ ] [50-Hashing with CBC-MAC](https://cryptopals.com/sets/7/challenges/50)
 * [ ] [51-Compression Ratio Side-Channel Attacks](https://cryptopals.com/sets/7/challenges/51)
 * [ ] [52-Iterated Hash Function Multicollisions](https://cryptopals.com/sets/7/challenges/52)
 * [ ] [53-Kelsey and Schneier's Expandable Messages](https://cryptopals.com/sets/7/challenges/53)
 * [ ] [54-Kelsey and Kohno's Nostradamus Attack](https://cryptopals.com/sets/7/challenges/54)
 * [ ] [55-MD4 Collisions](https://cryptopals.com/sets/7/challenges/55)
 * [ ] [56-RC4 Single-Byte Biases](https://cryptopals.com/sets/7/challenges/56)
- [ ] **Set 8 : Abstract Algebra**
 * [ ] 57-Diffie-Hellman Revisited: Small Subgroup Confinement
 * [ ] 58-Pollard's Method for Catching Kangaroos
 * [ ] 59-Elliptic Curve Diffie-Hellman and Invalid-Curve Attacks
 * [ ] 60-Single-Coordinate Ladders and Insecure Twists
 * [ ] 61-Duplicate-Signature Key Selection in ECDSA (and RSA)
 * [ ] 62-Key-Recovery Attacks on ECDSA with Biased Nonces
 * [ ] 63-Key-Recovery Attacks on GCM with Repeated Nonces
 * [ ] 64-Key-Recovery Attacks on GCM with a Truncated MAC
