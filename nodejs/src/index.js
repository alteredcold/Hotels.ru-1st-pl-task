function getComputerPhrase(count) {
    const lastNumber = count % 10;
    const lastTwoNumbers = count % 100;

    let word;
    switch (true) {
        case lastTwoNumbers >= 11 && lastTwoNumbers <= 14:
            word = "компьютеров";
            break;
        case lastNumber === 1:
            word = "компьютер";
            break;
        case lastNumber >= 2 && lastNumber <= 4:
            word = "компьютера";
            break;
        default:
            word = "компьютеров";
            break;
    }

    return `NodeJS: ${count} ${word}`;
}

console.log(getComputerPhrase(25));
console.log(getComputerPhrase(41));
console.log(getComputerPhrase(1048));
