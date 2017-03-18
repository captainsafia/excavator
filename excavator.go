package excavator

import "os"
import "os/user"
import "errors"
import "bufio"

func readLines(path string, numLines int) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    linesCounted := 0

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
        linesCounted += 1
        if linesCounted == numLines {
            break;
        }
    }

    return lines, scanner.Err()
}

func getHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        return ""
    }
    return usr.HomeDir
}

func getPathFromShell(shell string) (string, error) {
    switch shell {
        case "bash":
            return getHomeDir() + "/.bash_history", nil
        case "zsh":
            return getHomeDir() + "/.zsh_history", nil
        case "fish":
            return getHomeDir() + "/.local/share/fish/fish_history", nil
        default:
            return "", errors.New("Unsupported shell type!")
    }
}

func GetHistory(shell string, numCommands int) ([]string, error) {
    path, err := getPathFromShell(shell)
    if err != nil {
        return nil, err
    }

    commands, err := readLines(path, numCommands)
    if err != nil {
        return nil, err
    }
    return commands, err
}
