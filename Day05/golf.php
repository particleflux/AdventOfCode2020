<?for(;$p=fgets(STDIN);$c++)$i[]=base_convert(strtr($p,'BFRL','1010'),2,10);
sort($i);echo$i[--$c].' ';for(;$j++<$c;)if($i[$j+1]-$i[$j]>1)echo$i[$j]+1;
