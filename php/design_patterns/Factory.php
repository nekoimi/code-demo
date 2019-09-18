<?php
/**
 * # ------------Oooo---
 * # -----------(----)---
 * # ------------)--/----
 * # ------------(_/-
 * # ----oooO----
 * # ----(---)----
 * # -----\--(--
 * # ------\_)-
 * # ----
 * #     author : Yprisoner <yyprisoner@gmail.com>
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/

class People {
    public function say() { echo get_called_class() . PHP_EOL; }
}
class Student extends People {}
class Teacher extends People {}

interface FactoryInterface {
    public function Create(string $name);
}

class Factory implements FactoryInterface {

    public function create(string $name)
    {
        switch ($name) {
            case 'student':
                return new Student();
            case 'teacher':
                return new Teacher();
            default:
                throw new RuntimeException("Class not exists.");
        }
    }
}

$factory = new Factory();
$student = $factory->create('student');
$teacher = $factory->create('teacher');

$student->say();
$teacher->say();
/**
 * result:
 *
 * Student
 * Teacher
 */
