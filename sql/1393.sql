SELECT stock_name, sum(if(operation='sell', price,-price)) AS capital_gain_loss
FROM stocks
GROUP BY stock_name
