%- https://dota2.fandom.com/wiki/Helm_of_the_Overlord

:- discontiguous 
    parent/2, main_type/1, recipe_type/1, can_be_disassembled/2.

%- Item Y and be created from item X
parent('Helm of the Dominator', 'Helm of the Overlord').
parent('Vladmirs Offering', 'Helm of the Overlord').
parent('Helm of the Overlord Recipe', 'Helm of the Overlord').
parent('Helm of Iron Will', 'Helm of the Dominator').
parent('Diadem', 'Helm of the Dominator').
parent('Helm of the Dominator Recipe', 'Helm of the Dominator').
parent('Buckler', 'Vladmirs Offering').
parent('Ring of Basilius', 'Vladmirs Offering').
parent('Morbid Mask', 'Vladmirs Offering').
parent('Blades of Attack', 'Vladmirs Offering').
parent('Vladmirs Offering Recipe', 'Vladmirs Offering').
parent('Sages Mask', 'Ring of Basilius').
parent('Ring of Basilius Recipe', 'Ring of Basilius').
parent('Ring of Protection', 'Buckler').
parent('Buckler Recipe', 'Buckler').
parent('Morbid Mask', 'Mask of Madness').
parent('Quarterstaff', 'Mask of Madness').

%- Item X which is not a Recipe
main_type('Helm of the Overlord').
main_type('Helm of the Dominator').
main_type('Vladmirs Offering').
main_type('Helm of Iron Will').
main_type('Diadem').
main_type('Ring of Basilius').
main_type('Sages Mask').
main_type('Buckler').
main_type('Ring of Protection').
main_type('Morbid Mask').
main_type('Blades of Attack').
main_type('Mask of Madness').
main_type('Quarterstaff').
main_type('Tango').
main_type('Bottle').

%- Item X is a Recipe for another item.
recipe_type('Helm of the Overlord Recipe').
recipe_type('Helm of the Dominator Recipe').
recipe_type('Ring of Basilius Recipe').
recipe_type('Vladmirs Offering Recipe').
recipe_type('Buckler Recipe').

can_be_disassembled('Mask of Madness').

%- Item X must contain item Y to be created
child(X, Y) :-
    parent(Y, X).

%- Item Y includes item X on any of creation stage
can_be_assembled(X, Y) :-
    parent(X, Y); parent(X, Z), parent(Z, Y); parent(X, Z), parent(Z, A), parent(A, Y).

%- Item X includes item X on any of creation stage
contains_element(X, Y) :-
    can_be_assembled(Y, X).

%- You can disassemble item X to Y
can_be_dissasembled_into(Y, X) :-
    parent(X, Y), can_be_disassembled(Y);
    parent(X, Z), parent(Z, Y), can_be_disassembled(Y), can_be_disassembled(Z);
    parent(X, Z), parent(Z, A), parent(A, Y), can_be_disassembled(Y), can_be_disassembled(A), can_be_disassembled(Z).

%- Item X is recipe for item Y
recipe_for_item(X, Y) :-
    parent(X, Y), recipe_type(X).