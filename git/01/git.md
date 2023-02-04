* git init
//инициализация

git config --global user.name "<ваше_имя>"
git config --global user.email "<адрес_почты@email.com>"

rm -rf .git
удалить репозиторий

* git remote add origin git@github.com:igorsimdyanov/hello.git
подключиться к репозиторию

git status
git branch

//переключить ветку
git checkout имя ветки

//переключиться на предыдущий коммит
git checkout имя ветки^
git checkout HEAD^

//создать и переключиться на новую ветку
git checkout -b имя ветки

//добавить все изменения
git add .

//добавить коммит
git commit -m ''

git push -u origin master
git merge 

git branch -f main HEAD~3



git branch -D имя ветки 
удалить ветку

