load("varz", "set_varz", "get_varz")

def main(args):
    example, ok = get_varz("example")
    if ok:
        logf('get_varz("example")=%s', example)
    else:
        logf('get_varz("example")=unset')
    set_varz("example", str(time.now()))
    set_varz("list", [1,2,3])