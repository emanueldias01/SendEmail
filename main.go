package main

import (
    "bufio"
    "fmt"
    "net/smtp"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

var cmd *exec.Cmd

func clearTerminal() {
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("cmd", "/c", "cls")
    default:
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func sendEmail() error {
    clearTerminal()
    fmt.Println("WELCOME TO CLI TO SEND EMAIL BY GMAIL\n\n")

    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Write from:")
    from, _ := reader.ReadString('\n')
    from = strings.TrimSpace(from)

    fmt.Println("Write to:")
    to, _ := reader.ReadString('\n')
    to = strings.TrimSpace(to)

    fmt.Println("Write app password:")
    passwordApp, _ := reader.ReadString('\n')
    passwordApp = strings.TrimSpace(passwordApp)

    fmt.Println("Write a subject:")
    subject, _ := reader.ReadString('\n')
    subject = strings.TrimSpace(subject)

    fmt.Println("Write a body:")
    body, _ := reader.ReadString('\n')
    body = strings.TrimSpace(body)

    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    msg := []byte("To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n\r\n" +
        body + "\r\n")

    auth := smtp.PlainAuth("", from, passwordApp, smtpHost)

    clearTerminal()
    fmt.Println("Sending Email...")

    return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

func main() {
    for {
        err := sendEmail()
        if err != nil {
            fmt.Println("An error occurred while sending the email. Please try again.")
            fmt.Println("Press (y) to try again or any other key to exit:")
            var option string
            fmt.Scanln(&option)

            if strings.ToLower(option) != "y" {
                break
            }
        } else {
            fmt.Println("Email sent successfully!")
            break
        }
    }
}