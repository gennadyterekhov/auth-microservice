def main():
    removeVersionFromGoMod()


def filePutContents(fl: str, dt: str):
    flHandle = open(fl, 'w')
    flHandle.write(dt)
    flHandle.close()


def fileGetContents(fl: str) -> str:
    flHandle = open(fl, 'r')
    contents = flHandle.read()
    flHandle.close()
    return contents


# https://yandex.cloud/ru/docs/functions/lang/golang/dependencies
def removeVersionFromGoMod():
    print('removeVersionFromGoMod')
    filePutContents(
        'go.mod',
        fileGetContents('go.mod').replace('go 1.23.0', ''),
    )


if __name__ == '__main__':
    main()
