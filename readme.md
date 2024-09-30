

# 🎨 Ascii Art

## 🌟 Project Overview

Welcome to **Ascii Art**, where text meets creativity! 🖌️ This Go-powered command-line tool turns your boring text into eye-popping, stylized ASCII art. Whether you're looking to spice up your terminal output or create some retro-inspired text art, Ascii Art has got you covered! 🌈

## ✨ Features

- **🎨 Custom Font Magic**: Your text comes to life with a custom font defined in `standard.txt`. Each letter and symbol has its unique personality, just waiting to jump onto your screen!
- **📝 Multiline Marvel**: Want to create an entire masterpiece? Use `\\n` to craft stunning multiline ASCII art!
- **💻 Command-Line Wizardry**: Just type in your text and let the magic happen. Perfect for adding flair to your scripts or surprising your friends!

## 🚀 How It Works

1. **🔍 Font Loading**: We dive into the `standard.txt` file to fetch the artistic blueprints for your text. Each character gets its own 8-line high makeover.
   
2. **🧩 Input Processing**: Simply input your text at the command line. Want a new line? Just sprinkle a `\\n` in there, and voilà!

3. **🎭 ASCII Art Rendering**: We assemble your characters line by line, creating a dazzling display of ASCII artistry right in your terminal.

## 🎉 Usage

### 🛠️ Command-Line Fun

To unleash the magic, type:

```bash
go run main.go "YourTextHere"
```

- Replace `"YourTextHere"` with the text you want to dazzle the world with.
- Use `\\n` to create new lines in your masterpiece.

### 🌈 Example

```bash
go run main.go "Hello\\nWorld"
```

This command will light up your screen with:

```
✨ASCII art representation of "Hello"✨
✨ASCII art representation of "World"✨
```

## 🖼️ Font File (`standard.txt`)

The magic happens inside `standard.txt`, where each character's stylish ASCII look is carefully crafted across 9 lines. We calculate where each character’s art begins and ends, ensuring your text is always dressed to impress. ✨

## ⚠️ Error Handling

- If `standard.txt` decides to play hide and seek, we’ll let you know with a friendly error message and safely exit stage left. 🎭
- Got empty input? No worries! We’ll just print a nice, clean blank line for you.

## 🤝 Contributions

Got ideas to make Ascii Art even cooler? Fork this project and share your awesomeness with a pull request! 🚀

## 📜 License

Feel free to remix, share, and enjoy! This project is licensed under the [MIT License](LICENSE). 🎉

---