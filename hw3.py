import random

def random_dna(length):
    alphabet = ['A', 'T', 'G', 'C']
    width = 70
    seq = ''.join(random.choices(alphabet, k=length))
    return '\n'.join([seq[i:i+width] for i in range(0, len(seq), width)])

for i in [2, 3, 5]:
    f = open(f"random_dna{i}.txt", "w")
    f.write(random_dna(10**i))
    f.close()
