#try examples from https://pytestguide.readthedocs.io/en/latest/pytestGuide/index.html#code-to-test
from example_op import stat2Num
import pytest


def test_stat2Num():

    assert stat2Num(3, 2) == (5, 2.5)
    assert stat2Num(3, 2.0) == (5, 2.5)
    assert stat2Num(3.5, 2.5) == (6, 3)
    assert stat2Num(3.5, 2.5) == (6.0, 3.0)

    assert stat2Num(3.5, 2.5) == (6.0, 10.0)

"""
test_example_op.py F                                                                    [100%]

========================================== FAILURES ===========================================
________________________________________ test_stat2Num ________________________________________

    def test_stat2Num():

        assert stat2Num(3, 2) == (5, 2.5)
        assert stat2Num(3, 2.0) == (5, 2.5)
        assert stat2Num(3.5, 2.5) == (6, 3)
        assert stat2Num(3.5, 2.5) == (6.0, 3.0)

>       assert stat2Num(3.5, 2.5) == (6.0, 10.0)
E       assert (6.0, 3.0) == (6.0, 10.0)
E         At index 1 diff: 3.0 != 10.0
E         Use -v to get the full diff

test_example_op.py:12: AssertionError
=================================== short test summary info ===================================
FAILED test_example_op.py::test_stat2Num - assert (6.0, 3.0) == (6.0, 10.0)
====================================== 1 failed in 0.03s ======================================
"""



#tests can be grouped together
@pytest.mark.pyfile # marks can be used to run selected tests
class Test_stat2Num(object):

    def test_stat2Num_inputInt(self):
        assert stat2Num(3, 2) == (5, 2.5)
        assert stat2Num(3, 2) == (5.0, 2.5)

    def test_stat2Num_inputFloat(self):
        assert stat2Num(3, 2) == (5, 2.5)
        assert stat2Num(3, 2) == (5.0, 2.5)

"""
========================================== FAILURES ===========================================
________________________________________ test_stat2Num ________________________________________

    def test_stat2Num():

        assert stat2Num(3, 2) == (5, 2.5)
        assert stat2Num(3, 2.0) == (5, 2.5)
        assert stat2Num(3.5, 2.5) == (6, 3)
        assert stat2Num(3.5, 2.5) == (6.0, 3.0)

>       assert stat2Num(3.5, 2.5) == (6.0, 10.0)
E       assert (6.0, 3.0) == (6.0, 10.0)
E         At index 1 diff: 3.0 != 10.0
E         Use -v to get the full diff

test_example_op.py:12: AssertionError
=================================== short test summary info ===================================
FAILED test_example_op.py::test_stat2Num - assert (6.0, 3.0) == (6.0, 10.0)
================================= 1 failed, 2 passed in 0.04s =================================
"""