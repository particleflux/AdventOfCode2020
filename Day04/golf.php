<?$r=['byr):((19[2-9]\d|200[0-2]','iyr):((201\d|2020','eyr):((202\d|2030','hgt):(((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in','hcl):((#[\da-f]{6}','ecl):((amb|blu|brn|gry|grn|hzl|oth','pid):((\d{9}'];foreach(explode("

",file_get_contents('input'))as$p){for($s=$t=1,$i=0;$i<7;)if(!preg_match("/(?:^|\s)({$r[$i++]})(\s|$))?/",$p,$m))$s=$t=0;elseif(!($m[2]??0))$t=0;$x+=$s;$y+=$t;}echo"$x $y";
