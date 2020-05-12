# examples from https://docs.python.org/3/library/unittest.html
import unittest

"""
TestCase class provides assert methods to check for and report failures
some include:

assertEqual(a,b)
assertNotEqual(a,b)
assertTrue(x)
assertFalse(x)
assertIs(a,b)
assertIsNot(a,b)
assertIsNone(x)
assertIsNotNone(x)
assertIn(a,b)
assertNotIn(a,b)
assertIsInstance(a,b)
assertNotIsInstance(a,b)

assertRaises(exc, fun, *args, **kwds)
assertRaisesRegex(exc, r, fun, *args, **kwds)
assertWarns(warn, fun, *args, **kwds)
assertWarnsRegex(warn, r, fun, *args, **kwds)
assertLogs(logger, level)
"""

class TestStringMethods(unittest.TestCase):

    def test_upper(self): 
        self.assertEqual('foo'.upper(), 'FOO') # checks for an expected result

    def test_isupper(self):
        self.assertTrue('FOO'.isupper()) #verify a condition
        self.assertFalse('Foo'.isupper())

    def test_split(self):
        s = 'hello world'
        self.assertEqual(s.split(), [1, 'hello', 'world'])
        # check that s.split fails when the separator is not a string
        with self.assertRaises(TypeError): # verify that a specific exception gets raised
            s.split(2)

# setup method allows you to define instructions that will be executed before/after each test method


if __name__ == '__main__':
    unittest.main()


"""
.F.
======================================================================
FAIL: test_split (__main__.TestStringMethods)
----------------------------------------------------------------------
Traceback (most recent call last):
  File "unittest_example.py", line 15, in test_split
    self.assertEqual(s.split(), [1, 'hello', 'world'])
AssertionError: Lists differ: ['hello', 'world'] != [1, 'hello', 'world']

First differing element 0:
'hello'
1

Second list contains 1 additional elements.
First extra element 2:
'world'

- ['hello', 'world']
+ [1, 'hello', 'world']
?  +++

"""