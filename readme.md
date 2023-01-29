1. установить Xcode:
https://developer.apple.com/support/xcode/
Обновить Xcode
xcode-select --install
2. установить Homebrew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
3. открыть файл:
nano ~/.bash_profile
записать в файл:
export PATH=/usr/local/bin:$PATH
ctrl+o ctrl+x
активация настроек
source ~/.bash_profile
4. brew doctor
5. установить golang
brew install golang
brew update
brew upgrade golang
6. создать каталоги
mkdir -p $HOME/go/{bin,src}
открыть файл:
nano ~/.bash_profile
записать в файл:
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
ctrl+o ctrl+x
активация настроек
. ~/.bash_profile
7. инициализация git
git init
git config --global user.name "<ваше_имя>"
git config --global user.email "<адрес_почты@email.com>"
ssh-keygen -t ed25519 -C "your_email@example.com"
nano /Users/st/.ssh/id_ed25519.pub
git remote add origin git@github.com:StasTolmachov/Golang.git
git clone https://github.com/StasTolmachov/Golang.git
8. VS Code
9. Install oh-my-zsh:
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"



