#!/usr/bin/python3
'''creates a class'''


def lookup(obj):
    '''returns a list of attributes'''
    return dir(obj)


if __name__ == "__main__":
    class MyClass1(object):
        pass

    class MyClass2(object):
        '''seconde class'''
        my_attr1 = 3

        def my_meth(self):
            '''just a function definitiion'''
            pass

    print(lookup(MyClass1))
    print(lookup(MyClass2))
