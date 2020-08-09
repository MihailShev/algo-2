const pow = (a, n) => {
    let r = 1;

    for (let i = 0; i < n; i++) {
        r *= a;
    }

    console.log(r);
}

pow(1.001, 1000);

console.log(Math.pow(1.001, 1000));