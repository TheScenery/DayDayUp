function getAccountIndex(account) {
    return account - 1;
}

/**
 * @param {number[]} balance
 */
var Bank = function (balance) {
    this.balance = balance;
};

Bank.prototype.isValidAccount = function (accountIndex) {
    return accountIndex >= 0 && accountIndex < this.balance.length;
}

/** 
 * @param {number} account1 
 * @param {number} account2 
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.transfer = function (account1, account2, money) {
    const account1Index = getAccountIndex(account1);
    const account2Index = getAccountIndex(account2);
    if (!this.isValidAccount(account1Index) || !this.isValidAccount(account2Index)) {
        return false;
    }
    if (money > this.balance[account1Index]) {
        return false;
    }
    this.balance[account1Index] -= money;
    this.balance[account2Index] += money;
    return true;
};

/** 
 * @param {number} account 
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.deposit = function (account, money) {
    const accountIndex = getAccountIndex(account);
    if (!this.isValidAccount(accountIndex)) {
        return false;
    }
    this.balance[accountIndex] += money;
    return true;
};

/** 
 * @param {number} account 
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.withdraw = function (account, money) {
    const accountIndex = getAccountIndex(account);
    if (!this.isValidAccount(accountIndex)) {
        return false;
    }
    if (money > this.balance[accountIndex]) {
        return false;
    }
    this.balance[accountIndex] -= money;
    return true;
};

/**
 * Your Bank object will be instantiated and called as such:
 * var obj = new Bank(balance)
 * var param_1 = obj.transfer(account1,account2,money)
 * var param_2 = obj.deposit(account,money)
 * var param_3 = obj.withdraw(account,money)
 */