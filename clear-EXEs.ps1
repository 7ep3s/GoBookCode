get-childitem -Recurse *.exe | select -ExpandProperty fullname | foreach {remove-item $_ -Force}