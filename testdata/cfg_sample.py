def if_else(x):
    if x > 0:
        return "positive"
    else:
        return "non-positive"

def early_return(data):
    if not data:
        return None
    result = process(data)
    return result

def for_loop(items):
    result = ""
    for item in items:
        result = item
    return result

def nested_if_in_for(items):
    count = 0
    for item in items:
        if len(item) > 3:
            count = count + 1
    return count

def linear_function():
    x = 1
    y = x + 2
    return y

def empty_function():
    pass

def process(data):
    return data

def elif_no_else(x):
    result = 0
    if x > 10:
        result = 1
    elif x > 5:
        result = 2
    return result

def loop_with_break(items):
    for item in items:
        if item == "stop":
            break
    return "done"

def loop_with_continue(items):
    count = 0
    for item in items:
        if item == "":
            continue
        count = count + 1
    return count

def single_return():
    return 42
