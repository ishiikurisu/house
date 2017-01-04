import sys
import os
import house

def main():
    h = house.House(sys.argv[1:])
    h.draw()

if __name__ == '__main__':
    print('---')
    main()
    print('...')
