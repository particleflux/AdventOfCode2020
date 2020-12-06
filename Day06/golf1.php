<?for($i=0,$z=explode("

",file_get_contents('input'));$z[$i];){$a+=count($f=count_chars($z[$i++],1))-(($n=$f[10])>0);foreach($f as$k=>$v)$b+=$k!=10&$v==$n+1;}echo"$a $b";
