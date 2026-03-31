from string import ascii_lowercase

def is_pangram(sentence):
    return all(s in sentence.lower() for s in ascii_lowercase)

