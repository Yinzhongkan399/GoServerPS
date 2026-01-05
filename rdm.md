### Code 1
请给出一段代码，它的语言是{go},它的目的是，在本地开启一个http服务器，当用户向指定的URL发出GET或者POST请求时，根据URL调用对应的响应函数，并返回一个JSON文件。应当具有充分的可扩展性。

### Code 2
请给出一段代码，它的语言是{go},它完整重构文件夹下{Socket.py}的效果，ListAll函数对外部暴露，其余函数不对外部暴露，这段代码仅能在Linux上正确运行，所以不要进行运行测试。

### Code 3
请给出一段代码，它的语言是{go},它完整重构文件夹下{ReadBTFandGetItsMember.py}的效果，ReadBTFandGetItsMember函数对外部暴露，其余函数不对外部暴露，这段代码仅能在Linux上正确运行，所以不要进行运行测试。将生成的代码写入ReadBTFandGetItsMember.go这个文件，并请在README.md后给出这段代码的说明。


### Code 4
请给出一段代码，它的语言是{go},它完整重构文件夹下{ReadBTFandGetItsMember.go}的效果，ReadBTFandGetItsMember函数对外部暴露，其余函数不对外部暴露。重构的原因是，Go语言提供了bpf2go库，其中有对BTF功能的直接调用，这使得不需要人工读取BTF文件，也能获取到BTF信息，希望利用bpf2go的这个功能优势。
这段代码仅能在Linux上正确运行，所以不要进行运行测试。将生成的代码写入GoLangReadBTFandGetItsMember.go这个文件，并请在README.md后给出这段代码的说明。
（这个Prompt没有用，因为底层是btftool实现，而为了解耦合，这是必要的，还是btftool吧，挂docker里面）

### Code 5
请给出一段代码，它的语言是{go},它完整重构文件夹下{translateJSON.py}的效果，translateJSON函数对外部暴露，其余函数不对外部暴露，这段代码仅能在Linux上正确运行，所以不要进行运行测试。将生成的代码写入translateJSON.go这个文件，并请在README.md后给出这段代码的说明。

### Code 6
请给出一个Readme文件，名为{Dependency.md}，讲解如何在Ubuntu24.04上安装运行本项目的方法。

### Code 7
请给出一段代码，它的语言是{go},它完整重构文件夹下{baserun.py}的效果，包装成单个函数，函数对外部暴露，其余函数不对外部暴露。将生成的代码写入baserun.go这个文件，并请在README.md后给出这段代码的说明。

### Code 8
请给出一段代码，它的语言是{go},它依次运行BaseRun(),ReadBTFandGetItsMember()和TranslateJSON()，分别位于baserun.go,ReadBTFandGetItsMember.go和和translateJSON中,这是项目的主要功能模块。将生成的代码写入main.go这个文件，并请在README.md的最后给出这段代码的说明。

### Code 9
本项目中的主要函数已经完成，请为这个项目生成一个Makefile文件。并请在README.md的最后给出这段代码的说明。之后，生成一个Readme文件，告诉用户应该如何编译和运行本项目。
