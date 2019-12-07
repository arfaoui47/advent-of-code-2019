if __name__ == '__main__':
    print(sum(int(i) // 3 - 2 for i in open("input.txt").read().split('\n')))
