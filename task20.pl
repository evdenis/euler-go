#!/usr/bin/perl

use warnings;
use strict;


my $f = 1;
foreach (2 .. 100) {
   $f *= $_
}

my $sum = 0;
foreach (split(//, $f)) { $sum += $_ }

print $sum;

