# sovle-spelling-bee-fyne
A program for solving the New York Times spelling bee puzzle.
The program is written in go using fyne.io so it can run most anywhere.

# Running on Windows
## msys-x64
% go build -v main.go dict.go
% go run -v main.go dict.go


# Building .apk for Android Studio
## msys-x64 
% fyne package -os android -appID com.example.solvespellingbeefyne -icon assets/solve-spelling-bee-fyne.png
% fyne package -os windows -icon assets/solve-spelling-bee-fyne.png
% fyne package -os linux -icon assets/solve-spelling-bee-fyne.png
 
# Installing .apk to Android Studio
Drag Solve_Spelling_Bee.apk to emulator icon

# Installaion on Ubuntu 
tar xf 'Solve Spellin Bee.tar.xz'
sudo make install

## - Assumptions
 * GOHOME = /home/craigwarner
%sudo apt install libgl1-mesa-dev xorg-dev
%go run -v main.go dict.go

# Alternative Intall
1) sudo /home/craigwarner/go/bin/fyne get github.com/craig-warner/solve-spelling-bee-fyne


