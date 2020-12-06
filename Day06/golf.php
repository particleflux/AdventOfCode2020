<?foreach(explode("

",file_get_contents('input'))as$p){
$m=[];$n=substr_count($p,"
")+1;for($i=0;$c=$p[$i++]??0;)if($c!="
"){$m[$c]++;if($n==$m[$c])$b++;}$a+=count($m);}echo"$a $b";