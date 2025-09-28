$min = 10000;

@list = (20, -20, 10, 5 ,35);

foreach $b (@list) {
	if($min > $b) {
		$min = $b;
	}
}
printf "min in (20 -20 10 5 35) = $min\n";