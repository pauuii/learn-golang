# 标准库参考地址
1. https://golang.google.cn/pkg/

# 标准库说明
1. archive/tar: 实现对tar存档的访问
2. archive/zip: 支持读取和写入zip存档
3. arena: 提供了为Go值集合分配内存的能力，并可以一次安全地手动释放这些空间
4. bufio: 实现缓冲I/O。包装io.Reader或io.Writer对象，创建另一个对象
5. builtin: 提供了Go预定义标识符的文档
6. bytes: 实现字节片操作的函数
7. compress/bzip2: 	实现了bzip2解压缩
8. compress/flate: 实现RFC 1951中描述的DEFLATE压缩数据格式
9. compress/gzip:  按照RFC 1952的规定，实现gzip格式压缩文件的读写
10. compress/lzw: 实现了Lempel-Ziv-Welch压缩数据格式
11. compress/zlib:  实现zlib格式压缩数据的读写，按RFC 1950中所规定
12. container/heap: 实现heap.Interface的任何类型提供堆操作
13. container/list: 实现了一个双链接列表
14. container/ring: 实现了环形链表的操作
15. context: (上下文)定义Context类型，它跨API边界和进程之间携带截止日期、取消信号和其他请求范围值
16. crypto: 收集常用的加密常数
17. crypto/aes: 实施AES加密（以前称为Rijndael）
18. crypto/boring: 公开仅在使用Go+BoringCrypto构建时可用的函数
19. crypto/cipher: 实现了标准的分组密码模式，这些模式可以封装在低级分组密码实现中
20. crypto/des: 实现数据加密标准（DES）和三重数据加密算法（TDEA）
21. crypto/dsa: 实现了FIPS 186-3中定义的数字签名算法
22. crypto/ecdh: 在NIST曲线和Curve25519上实现椭圆曲线Diffie-Hellman
23. crypto/ecdsa: 实现了FIPS 186-4和SEC 1 2.0版中定义的椭圆曲线数字签名算法
24. crypto/ed25519: 实现Ed25519签名算法
25. crypto/elliptic: 在素数字段上实现实现标准NIST P-224、P-256、P-384和P-521椭圆曲线
26. crypto/hmac: 实现了密钥散列消息认证码（hmac）
27. crypto/md5: 实现RFC 1321中定义的MD5哈希算法
28. crypto/rand: 实现密码安全的随机数生成器
29. crypto/rc4: 实现了Bruce Schneier的应用密码学中定义的RC4加密
30. crypto/rsa: 按照PKCS#1和RFC 8017中的规定实现RSA加密
31. crypto/sha1: 实现了RFC 3174中定义的SHA-1哈希算法
32. crypto/sha256: 实现了FIPS 180-4中定义的SHA224和SHA256散列算法
33. crypto/sha512: 实现FIPS 180-4中定义的SHA-384、SHA-512、SHA-512/224和SHA-512/256哈希算法
34. crypto/subtle: 在密码代码中通常有用但需要仔细考虑才能正确使用的函数
35. crypto/tls: 部分实现了RFC 5246中规定的TLS 1.2和RFC 8446中指定的TLS 1.3
36. crypto/tls/fipsonly: 仅将所有TLS配置限制为FIPS批准的设置
37. crypto/x509: 实现X.509标准的子集
38. crypto/pkix: 包含用于ASN.1解析和序列化X.509证书、CRL和OCSP的共享低级结构
39. database/sql: 围绕SQL（或类似SQL）数据库提供通用接口
40. database/sql/driver: 定义要由包sql使用的数据库驱动程序实现的接口
41. debug/buildinfo: buildinfo提供了对嵌入在Go二进制文件中的信息的访问，这些信息是如何构建的
42. debug/dwarf: dwarf提供对从可执行文件加载的dwarf调试信息的访问，如dwarf 2.0标准中所定义
43. debug/elf: elf实现对elf对象文件的访问
44. debug/gosym: gosym实现了对gc编译器生成的Go二进制文件中嵌入的Go符号和行号表的访问
45. debug/macho: macho实现对Mach-O对象文件的访问
46. debug/pe: pe实现对pe（Microsoft Windows可移植可执行文件）文件的访问
47. debug/plan9obj: plan9obj实现对Plan9 a.out对象文件的访问
48. embed: 提供对运行的Go程序中嵌入的文件的访问
49. encoding/ascii85: 实现btoa工具和Adobe的PostScript和PDF文档格式中使用的ascii85数据编码
50. encoding/asn1: 按照ITU-T Rec X.690的定义，实现对DER编码ASN.1数据结构的解析
51. encoding/base32: 实现RFC 4648指定的base32编码
52. encoding/base64: 实现RFC 4648指定的base64编码
53. encoding/binary: 实现数字和字节序列之间的简单转换以及变量的编码和解码
54. encoding/csv: 读取和写入逗号分隔值（CSV）文件
55. encoding/gob: 管理gobs流-编码器（发射器）和解码器（接收器）之间交换的二进制值
56. encoding/hex: 实现十六进制编码和解码
57. encoding/json: 实现RFC 7159中定义的JSON编码和解码
58. encoding/pem: 实现PEM数据编码，该编码源自隐私增强邮件
59. encoding/xml: 实现一个简单的XML1.0解析器，它可以理解XML名称空间
60. errors: 实现处理错误的函数
61. expvar: 为公共变量（如服务器中的操作计数器）提供标准化接口
62. flag: 实现命令行标志解析
63. fmt: 使用类似于C的printf和scanf的函数实现格式化的I/O
64. go/ast: 声明用于表示Go包的语法树的类型
65. go/build: 收集有关Go软件包的信息
66. go/build/constraint: 实现构建约束行的解析和求值
67. go/constant: 实现表示非类型化Go常量及其相应操作的值
68. go/doc: 从Go AST中提取源代码文档
69. go/comment: 实现Go文档注释（文档注释）的解析和重新格式化，这些注释直接位于包、const、func、type或var的顶级声明之前
70. go/format: 实现Go源的标准格式
71. go/importer: 提供对导出数据导入器的访问
72. go/parser: 实现Go源文件的解析器
73. go/printer: 实现AST节点的打印
74. go/scanner: 实现Go源文本的扫描程序
75. go/token: 定义表示Go编程语言的词汇标记的常量和标记的基本操作（打印、谓词）
76. go/types: 声明数据类型并实现Go包的类型检查算法
77. hash: 为哈希函数提供接口
78. hash/adler32: 实现Adler-32校验和
79. hash/crc32: 实现32位循环冗余校验或CRC-32校验和
80. hash/crc64: 实现64位循环冗余校验或CRC-64校验和
81. hash/fnv: 实现了FNV-1和FNV-1a，由Glenn Fowler、Landon Curt Noll和Phong Vo创建的非加密散列函数
82. hash/maphash: 提供字节序列的哈希函数
83. html: 提供了转义和取消HTML文本的功能
84. html/template: 实现了数据驱动的模板，用于生成HTML输出，防止代码注入
85. image: 实现基本的二维图像库
86. image/color: 实现基本颜色库
87. image/color/palette: 提供标准调色板
88. image/draw: 提供图像合成功能
89. image/gif: 实现GIF图像解码器和编码器
90. image/jpeg: 实现JPEG图像解码器和编码器
91. image/png: 实现PNG图像解码器和编码器
92. index/suffixarray: 使用内存中的后缀数组在对数时间内实现子字符串搜索
93. io: 提供I/O原语的基本接口
94. io/fs: 定义文件系统的基本接口
95. ioutil: 实现一些I/O实用程序功能
96. log: 实现一个简单的日志记录包
97. log/syslog: 提供系统日志服务的简单界面
98. math: 提供基本常数和数学函数
99. math/big: 实现任意精度算术（大数字）
100. math/bits: 为预先声明的无符号整数类型实现位计数和操作函数
101. math/cmplx: 为复数提供基本常数和数学函数
102. math/rand: 实现不适合安全敏感工作的伪随机数生成器
103. mime: 实现MIME规范的部分
104. mime/multipart: 实现MIME多部分解析，如RFC 2046中所定义
105. mime/quotedprintable: 实现RFC 2045指定的引用可打印编码
106. net: 为网络I/O提供可移植接口，包括TCP/IP、UDP、域名解析和Unix域套接字
107. net/http: 提供HTTP客户端和服务器实现
108. net/http/cgi: 实现RFC 3875中规定的CGI（公共网关接口）
109. net/http/cookiejar: 实现符合内存RFC 6265的http.CookieJar
110. net/http/fcgi: 实现FastCGI协议
111. net/http/httptest: 提供用于HTTP测试的实用程序
112. net/http/httptrace: 提供跟踪HTTP客户端请求中事件的机制
113. net/http/httputil: 提供了HTTP实用程序功能，补充了net/HTTP包中更常见的功能
114. net/http/pprof: 通过其HTTP服务器运行时分析数据，以pprof可视化工具预期的格式提供服务
115. net/mail: 实现邮件消息的解析
116. net/netip: 定义一个小值类型的IP地址类型
117. net/rpc: 通过网络或其他I/O连接访问对象的导出方法
118. net/rpc/jsonrpc: 为RPC包实现JSON-RPC 1.0 ClientCodec和ServerCodec
119. net/smtp: 实现RFC 5321中定义的简单邮件传输协议
120. net/textproto: 以HTTP、NNTP和SMTP的形式实现对基于文本的请求/响应协议的通用支持
121. net/url: 解析URL并实现查询转义
122. os: 为操作系统功能提供独立于平台的接口
123. os/exec: 运行外部命令
124. os/signal: 实现对输入信号的访问
125. os/user: 允许按名称或id查找用户帐户
126. path: 实现用于操作斜杠分隔路径的实用程序例程
127. path/filepath: 实现实用程序例程，以与目标操作系统定义的文件路径兼容的方式操作文件名路径
128. plugin: 实现Go插件的加载和符号解析
129. reflect: 实现运行时反射，允许程序操作具有任意类型的对象
130. regexp: 实现正则表达式搜索
131. regexp/syntax: 将正则表达式解析为解析树，并将解析树编译为程序
132. runtime: 包含与Go的运行时系统交互的操作，例如控制goroutine的函数
133. runtime/cgo: 包含对cgo工具生成的代码的运行时支持
134. runtime/coverage: 覆盖率数据操作
135. runtime/debug: 包含程序运行时自行调试的工具
136. runtime/metrics: 提供了一个稳定的接口来访问Go运行时导出的实现定义的度量
137. runtime/pprof: 以pprof可视化工具预期的格式写入运行时分析数据
138. runtime/race: 实现数据竞争检测逻辑
139. runtime/trace: 包含程序生成Go执行跟踪程序跟踪的工具
140. sort: 提供用于排序切片和用户定义集合的原语
141. strconv: 实现与基本数据类型的字符串表示之间的转换
142. strings: 实现操作UTF-8编码字符串的简单函数
143. sync: 提供基本的同步原语，如互斥锁
144. sync/atomic: 提供用于实现同步算法的低级原子内存原语
145. syscall: 包含到低级操作系统原语的接口
146. syscall/js: 使用js/wasm体系结构时，可以访问WebAssembly主机环境
147. testing: 为Go包的自动测试提供支持
148. testing/fstest: 实现对文件系统的测试实现和用户的支持
149. testing/iotest: 实现主要用于测试的读写器
150. testing/quick: 实现实用程序函数以帮助进行黑盒测试
151. text/scanner: 为UTF-8编码文本提供扫描仪和标记器
152. text/tabwriter: 实现写入筛选器（tabwriter.Writer），将输入中的选项卡列转换为正确对齐的文本
153. text/template: 实现用于生成文本输出的数据驱动模板
154. text/template/parse: 为text/template和html/template定义的模板构建解析树
155. time: 提供测量和显示时间的功能
156. time/tzdata: 提供时区数据库的嵌入副本
157. unicode: 提供数据和函数来测试Unicode代码点的某些属性
158. unicode/utf16: 实现UTF-16序列的编码和解码
159. unicode/utf8: 实现函数和常量以支持以UTF-8编码的文本
160. unsafe: 包含绕过Go程序类型安全的操作

# 常用包
1. context
2. encoding/json
3. encoding/base64
4. encoding/hex
5. errors
6. flag
7. fmt
8. io
9. io/util
10. log
11. math
12. net
13. net/http
14. net/url
15. reflect
16. regexp
17. sort
18. strconv
19. strings
20. sync
21. time

# 案例演示
1. encoding/json
2. encoding/base64
3. encoding/hex
4. errors
5. fmt
6. log
7. math
8. reflect
9. regexp
10. sort
