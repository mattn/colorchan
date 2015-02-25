# colorchan

コロちゃん

## Before

```
C:\>lein cljfmt check
```
![](http://go-gyazo.appspot.com/bb255a7034b463ff.png)

So Bad!

## After

```
C:\>lein cljfmt check 2>&1 | colorchan
```
![](http://go-gyazo.appspot.com/7333df50292d1e2b.png)

So Good!

## Usage

### Pipe
```
C:\>super-colorfull-command 2>&1 | colorchan
```

### Command
```
C:\>colorchan super-colorfull-command -great-flag
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
