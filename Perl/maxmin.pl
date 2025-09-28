$max = 1;
$min = 100;

@list = (10, 20, 30, 40 ,50);

foreach $a (@list) {
	if($max < $a) {
		$max = $a;
	}
}
printf "max = $max\n";

foreach $b (@list) {
	if($min > $b) {
		$min = $b;
	}
}
printf "min = $min\n";