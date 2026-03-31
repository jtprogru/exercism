def largest_product(series, size):
    if size < 0:
        raise ValueError('span must not be negative')
    if size > len(series):
        raise ValueError('span must be smaller than string length')
    if not all(c.isdigit() for c in series):
        raise ValueError('digits input must only contain digits')
    if size == 0:
        return 1

    digits = [int(c) for c in series]
    max_product = 0
    for i in range(len(digits) - size + 1):
        product = 1
        for j in range(i, i + size):
            product *= digits[j]
        max_product = max(max_product, product)
    return max_product