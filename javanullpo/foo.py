

class Foo:
    def is_null(self):
        return self is None


f: Foo = Foo()
print(f.is_null())  # False


x: Foo = None
print(x.is_null())
# Traceback (most recent call last):
#   File "foo.py", line 12, in <module>
#     print(x.is_null())
# AttributeError: 'NoneType' object has no attribute 'is_null'
